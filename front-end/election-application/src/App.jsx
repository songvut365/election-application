import react, {useState, useEffect} from 'react';
import axios from 'axios';
import Candidate from './components/Candidate';

function App() {
  const CANDIDATE_URL = import.meta.env.VITE_CANDIDATE_URL
  const ELECTION_URL = import.meta.env.VITE_ELECTION_URL

  const [candidates, setCandidates] = useState([])
  const [electionStatus, setElectionStatus] = useState({})

  useEffect(() => {
    const fetchCandidates = async () => {
      const response = await axios(CANDIDATE_URL);
      setCandidates(response.data);
    };
    fetchCandidates();

    const fetchElectionStatus = async () => {
      const response = await axios(ELECTION_URL+"/status");
      setElectionStatus(response.data);
    };
    fetchElectionStatus();
  }, []);

  return (
    <div className="container mx-auto text-center py-8">
      <h1 className='text-6xl mb-8 text-blue-700 font-bold'>Election Application</h1>

      <div className='grid gap-8 grid-cols-1 lg:grid-cols-3 xl:grid-cols-4'>
        {candidates.map((candidate) => (
          <Candidate candidate={candidate} key={candidate.ID} electionStatus={electionStatus.enable} />
        ))}
      </div>
    </div>
  )
}

export default App
