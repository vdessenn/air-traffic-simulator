import { useState } from "react";
import './SettingsPane.css'
import "leaflet/dist/leaflet.css";
import { setSimulationSpeed } from '../services/api';

export default function SettingsPane({setShowBoundaries, setShowAirports, setShowRisks}) {
  const [executionSpeedValue, setExecutionSpeedValue] = useState(1);

  return (
    <div className="settings-pane">
        <div id="speed-settings-container" className="column-settings-pane">
            <label htmlFor="volume"><b>Vitesse d'exécution :</b></label>
            <span>{executionSpeedValue}</span>
            <input  type="range"
                    id="execution-speed"
                    value={executionSpeedValue}
                    name="executionSpeed"
                    min="0.1"
                    max="10.0" step="0.1"
                    onChange={(e) => {
                    setSimulationSpeed(Number(e.target.value))
                    setExecutionSpeedValue(Number(e.target.value))
                    }} />
        </div>
        <div id="filters-container" className="column-settings-pane">
          <div className="filter-row">
            <input type="checkbox"
                   id="show-boundaries"
                   name="showBoundaries"
                   onChange={(e) => {
                    setShowBoundaries(e.target.checked)
                   }} />
            <label htmlFor="show-boundaries">Afficher les frontières</label>
          </div>
          <div className="filter-row">
            <input type="checkbox"
                   id="show-airports"
                   name="showAirports"
                   onChange={(e) => {
                    setShowAirports(e.target.checked)
                   }} />
            <label htmlFor="show-airports">Afficher les aéroports</label>
          </div>
          <div className="filter-row">
            <input type="checkbox"
                   id="show-risks"
                   name="showRisks"
                   onChange={(e) => {
                    setShowRisks(e.target.checked)
                   }} />
            <label htmlFor="show-risks">Afficher les conflits</label>
          </div>
        </div>

    </div>
  )
}