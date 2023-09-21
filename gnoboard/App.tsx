import { useGno } from '@gno/hooks/use-gno';
import Loading from '@gno/screens/loading';
import CustomRouter from '@gno/router/custom-router';
import { useEffect, useState } from 'react';

export default function App() {
  const gno = useGno();
  const [loading, setLoading] = useState<string | undefined>(undefined);

  useEffect(() => {
    setLoading('Initializing bridge...');
    gno
      .initBridge()
      .catch((error) => console.log(error))
      .finally(() => setLoading(undefined));
  }, []);

  if (loading) {
    return <Loading message={loading} />;
  }

  return <CustomRouter />;
}
