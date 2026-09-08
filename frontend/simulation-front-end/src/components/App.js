import './App.css'
import { useEffect, useState } from "react";
import useFetch from '../hooks/useFetch';
import { fetchBoundaries, fetchPlaneData, fetchAirports, fetchRisksSummary, fetchRisks } from '../services/api';
import MapView from './MapView';
import LoadingView from './LoadingView';
import ErrorView from './ErrorView';


function App() {
  const [tick, setTick] = useState(0);
  const { data: planeData, loading, error } = useFetch(fetchPlaneData, [tick]);
  const { data: boundariesData } = useFetch(fetchBoundaries);
  const { data: airportsData } = useFetch(fetchAirports);
  const { data: risksSummaryData } = useFetch(fetchRisksSummary, [tick]);
  const { data: risksData } = useFetch(fetchRisks, [tick]);

  useEffect(() => {
    const interval = setInterval(() => {
      setTick((t) => t + 1);
    }, 2000);

    return () => clearInterval(interval);
  }, []);
  if (loading) return <LoadingView />;
  if (error) return <ErrorView />;

  return <div className="App"><MapView planeData={planeData} boundariesData={boundariesData} airportsData={airportsData} risksData={risksData} risksSummaryData={risksSummaryData}/></div>;
}

export default App;
