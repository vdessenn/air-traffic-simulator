import { useState } from "react";
import './MetricsPane.css'
import "leaflet/dist/leaflet.css";
import warningImage from "../assets/warning.png";
import emergencyImage from "../assets/emergency.png";
import collisionImage from "../assets/crash.png";

export default function MetricsPane({risksSummaryData}) {
  // Provide default values if data is not loaded yet
  const safetyNotAssured = risksSummaryData?.safetyNotAssured ?? 0;
  const riskOfCollision = risksSummaryData?.riskOfCollision ?? 0;
  const collision = risksSummaryData?.collision ?? 0;

  return (
    <div className="metrics-overview">
      <div>
        <label><b>Risques en cours</b></label>
        <div id="risks-data-container">
          <div className="risk-category">
              <img src={warningImage} alt="Warning" />
              { safetyNotAssured }
          </div>

          <div className="risk-category">
              <img src={emergencyImage} alt="Emergency" />
              { riskOfCollision }
          </div>

          <div className="risk-category">
              <img src={collisionImage} alt="Collision" />
              { collision }
          </div>
        </div>
      </div>
    </div>
  )
}