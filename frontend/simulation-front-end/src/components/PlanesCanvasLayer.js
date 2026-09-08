import { useMap } from 'react-leaflet';
import { useEffect, useRef, useState } from 'react';
import L from 'leaflet';
import planeImg from '../assets/plane.png';
import RouteView from './RouteView';

export default function PlanesCanvasLayer({ planeData, activePlane, setActivePlane }) {
    const map = useMap();
    const canvasRef = useRef(null);
    const [image, setImage] = useState(null);

    // Load image once
    useEffect(() => {
        const img = new Image();
        img.src = planeImg;
        img.onload = () => setImage(img);
    }, []);

    const draw = () => {
        const canvas = canvasRef.current;
        if (!canvas || !map || !image || !planeData) return;

        const ctx = canvas.getContext('2d');
        const size = map.getSize();

        // Resize canvas if needed
        if (canvas.width !== size.x || canvas.height !== size.y) {
            canvas.width = size.x;
            canvas.height = size.y;
            canvas.style.width = size.x + 'px';
            canvas.style.height = size.y + 'px';
        } else {
            ctx.clearRect(0, 0, canvas.width, canvas.height);
        }

        const bounds = map.getBounds();

        planeData.forEach(plane => {
            // Simple culling
            if (!bounds.contains([plane.coordinates.latitude, plane.coordinates.longitude])) return;
            const point = map.latLngToContainerPoint([plane.coordinates.latitude, plane.coordinates.longitude]);

            ctx.save();
            ctx.translate(point.x, point.y);
            ctx.rotate(plane.direction * Math.PI / 180);
            ctx.drawImage(image, -12, -12, 24, 24); // Center the 24x24 icon
            ctx.restore();
        });
    };

    useEffect(() => {
        if (!map) return;

        // Create canvas
        const canvas = L.DomUtil.create('canvas', 'leaflet-zoom-animated');
        canvas.style.zIndex = 0;
        canvasRef.current = canvas;

        const pane = map.getPane('overlayPane');
        pane.appendChild(canvas);

        return () => {
            pane.removeChild(canvas);
        };
    }, [map, image]); // Re-bind if map changes (unlikely) or image loads

    // Redraw when data changes
    useEffect(() => {
        const canvas = canvasRef.current;
        const onMapUpdate = () => {
            const topLeft = map.containerPointToLayerPoint([0, 0]);
            L.DomUtil.setPosition(canvas, topLeft);
            draw();
        }

        const onZoomAnim = (e) => {
            const scale = map.getZoomScale(e.zoom);
            const offset = map._latLngToNewLayerPoint(map.containerPointToLatLng([0, 0]), e.zoom, e.center);

            L.DomUtil.setTransform(canvas, offset, scale);
        };

        map.on('move', onMapUpdate);
        map.on('moveend', onMapUpdate);
        map.on('zoomend', onMapUpdate);
        map.on('zoomanim', onZoomAnim);
        map.on('viewreset', onMapUpdate);

        // Initial draw
        draw();

        return () => {
            map.off('move', onMapUpdate);
            map.off('moveend', onMapUpdate);
            map.off('zoomend', onMapUpdate);
            map.off('zoomanim', onZoomAnim);
            map.off('viewreset', onMapUpdate);
        };
    });

    // Handle clicks
    useEffect(() => {
        const canvas = canvasRef.current;
        if (!canvas) return;

        const onClick = (e) => {
            if (!planeData) return;

            // Calculate click position relative to canvas
            const rect = canvas.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;

            // Find clicked plane (reverse order to find top-most first)
            let clicked = null;
            for (let i = planeData.length - 1; i >= 0; i--) {
                const plane = planeData[i];
                const point = map.latLngToContainerPoint([plane.coordinates.latitude, plane.coordinates.longitude]);

                // Distance check (12px radius)
                const dx = x - point.x;
                const dy = y - point.y;
                if (dx * dx + dy * dy <= 144) { // 12^2
                    clicked = plane;
                    break;
                }
            }

            if (clicked) {
                if (activePlane === clicked.id) {
                    setActivePlane("");
                } else {
                    setActivePlane(clicked.id);
                }
            }
        };

        canvas.addEventListener('click', onClick);
        return () => canvas.removeEventListener('click', onClick);
    }, [map, planeData, activePlane, setActivePlane]);

    // Find active plane object
    const activePlaneObj = planeData?.find(p => p.id === activePlane);

    return (
        <>
            {activePlaneObj &&
                <RouteView route={[[activePlaneObj.coordinates.latitude, activePlaneObj.coordinates.longitude]].concat(activePlaneObj.flightPlan)} />
            }
        </>
    );
}
