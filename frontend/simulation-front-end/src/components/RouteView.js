import { CircleMarker, Polyline } from 'react-leaflet'


export default function RouteView({route}) {
    return (
        <div>
            {
                route.map(tag =>
                    <CircleMarker center={tag} pathOptions={{ color: "orange" }} radius={2}/>
                )
            }
            <Polyline pathOptions={{ color: "orange" }} positions={route} />
        </div>
    )
}
