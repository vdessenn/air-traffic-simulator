import { useMap } from "react-leaflet";
import { useEffect } from "react";

export default function FixMapResize() {
  const map = useMap();

  useEffect(() => {
    setTimeout(() => {
      map.invalidateSize();
    }, 180);
  }, [map]);

  return null;
}