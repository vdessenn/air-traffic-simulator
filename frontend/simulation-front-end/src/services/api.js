export async function fetchPlaneData() {
  const res = await fetch("http://localhost:7500/planes");
  if (!res.ok) {
    throw new Error("Failed to fetch plane data");
  }

  var jsonValue = await res.json()
  var filledJson = []
  jsonValue.forEach(element => {
    let rawPlan = element.FlightPlan || [];
    let tags = rawPlan.map(wp => {
      return [wp.Coordinate.Latitude, wp.Coordinate.Longitude];
    }).filter(coords => coords !== undefined);
    filledJson.push(
        {
          id: element.Callsign,
          coordinates: {
            latitude: element.Position.Latitude,
            longitude: element.Position.Longitude
          },
          altitude: element.Position.Altitude,
          direction: element.Heading,
          flightPlan: tags
        }
    )

  })

  return filledJson;
}

async function fetchTagsData() {
  const res = await fetch("http://localhost:7500/tags");
  if (!res.ok) {
    throw new Error("Failed to fetch tags data");
  }
  const jsonValue = await res.json()

  var dictionary = {}
  jsonValue.forEach(element => {
    dictionary[element.id] = [element.coordinates.latitude, element.coordinates.longitude]
  });

  return dictionary;
}

export async function fetchRisks() {
  const res = await fetch("http://localhost:7500/risks");
  if (!res.ok) {
    throw new Error("Failed to fetch tags data");
  }
  const jsonValue = await res.json()

  var array = []
  jsonValue.forEach(element => {
    array.push({
      id: element.id,
      level: element.level,
      message: element.message,
      coordinates: {
        lat: element.location.latitude,
        lon: element.location.longitude
      }
    })
  }
  );
  console.log(array)
  return array;
}

export async function fetchBoundaries() {
  const res = await fetch("http://localhost:7500/data/airspace.json");
  if (!res.ok) {
    throw new Error("Failed to fetch tags data");
  }
  const jsonValue = await res.json()

  var regionList = jsonValue.gbr.filter(region => region.GbrUid.TxtName === "ZR:LF").map(region => {
    return region.Gbv.map(point => {
      return [point.GeoLat, point.GeoLong]
    })
  });
  return regionList;
}

export async function fetchAirports() {
  const res = await fetch("http://localhost:7500/data/airport.json");
  if (!res.ok) {
    throw new Error("Failed to fetch airports data");
  }
  const jsonValue = await res.json()

  var airportList = jsonValue.ahp.filter(airport => airport.CodeIcao.substring(0, 2) === "LF" && airport.CodeType === "AD").map(airport => {
    return {"name": airport.TxtName, "coordinates": [airport.GeoLat, airport.GeoLong]}
  });
  return airportList;
}

export async function setSimulationSpeed(speed) {
  await fetch("http://localhost:7500/simulationSpeed", {
    method: "POST",
    headers: {
      "Content-Type": "text/plain"
    },
    body: JSON.stringify(speed)
  });
}

export async function fetchRisksSummary() {
  const res = await fetch("http://localhost:7500/risks/summary");
  if (!res.ok) {
    throw new Error("Failed to fetch risks summary data");
  }
  const jsonValue = await res.json()

  return {
    safetyNotAssured: jsonValue[0],
    riskOfCollision: jsonValue[1],
    collision: jsonValue[2],
  }
}