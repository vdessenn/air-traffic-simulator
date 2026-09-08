import PlaneIcon from "./planeIcon.js";
import { CircleMarker, Popup } from 'react-leaflet'

export default function RiskMarker({ risk }) {
    const style = { color: risk.level == "COLLISION" ? "red" : (risk.level == "WARNING" ? "yellow" : "orange") }
    return (
        <CircleMarker center={risk.coordinates} pathOptions={style} radius={6}>
            <Popup>{risk.id}</Popup>
        </CircleMarker>
    )
}