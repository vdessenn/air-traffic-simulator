import { MapContainer, TileLayer, Polyline, CircleMarker, Popup } from 'react-leaflet'
import { useState } from "react";
import './MapView.css'
import "leaflet/dist/leaflet.css";
import PlanesCanvasLayer from './PlanesCanvasLayer';
import FixMapResize from '../utils/FixMapResize'
import SettingsPane from './SettingsPane';
import MetricsPane from './MetricsPane';
import RiskMarker from './RiskMarker';

const franceCenterCoordinates = [46.5119027, 2.9088057]

export default function MapView({ planeData, boundariesData, airportsData, risksData, risksSummaryData }) {
  const [activePlane, setActivePlane] = useState("");
  const [showBoundaries, setShowBoundaries] = useState(false);
  const [showAirports, setShowAirports] = useState(false);
  const [showRisks, setShowRisks] = useState(false);

  return (
    <div>
      <MapContainer center={franceCenterCoordinates} zoom={5} maxBounds={[[51.1635653, -6.0411473], [41.268622, 10.0936867]]} minZoom={5} scrollWheelZoom={true}>
        <FixMapResize />
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors | SIA AIRAC 12/25'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />

        <PlanesCanvasLayer
          planeData={planeData || []}
          activePlane={activePlane}
          setActivePlane={setActivePlane}
        />

        {
          showBoundaries &&
          boundariesData?.map(region => (
            <Polyline pathOptions={{ color: "red" }} positions={region} />
          ))
        }
        {
          showAirports &&
          airportsData?.map(airport => (
            <CircleMarker center={airport.coordinates} pathOptions={{ color: "blue" }} radius={2}>
              <Popup>{airport.name}</Popup>
            </CircleMarker>
          ))
        }
        {
          showRisks &&
          risksData?.map(risk => (
            <RiskMarker risk={risk} />
            
          ))
        }
      </MapContainer>
      <SettingsPane setShowBoundaries={setShowBoundaries} setShowAirports={setShowAirports} setShowRisks={setShowRisks} />
      <MetricsPane risksSummaryData={risksSummaryData} />
    </div>)
}