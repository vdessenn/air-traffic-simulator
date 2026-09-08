import L from "leaflet";
import planeImg from "../assets/plane.png";

export default function PlaneIcon(direction = 0) {
  return L.divIcon({
    className: "plane-icon",
    html: `
      <img 
        src="${planeImg}" 
        style="
          width: 24px;
          height: 24px;
          transform: rotate(${direction}deg);
          transform-origin: center;
        "
      />
    `,
    iconSize: [24, 24],
    iconAnchor: [12, 12],
  });
}