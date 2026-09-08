import { useEffect, useState } from "react";

export default function useFetch(asyncFunction, deps = []) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    let mounted = true;

    asyncFunction()
      .then((result) => mounted && setData(result))
      .catch((err) => mounted && setError(err))
      .finally(() => mounted && setLoading(false));

    return () => { mounted = false };
  }, deps);

  return { data, loading, error };
}
