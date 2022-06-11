import react, {useState, useEffect} from 'react';
import axios from 'axios';

import Candidate from './components/Candidate';
import VoteBox from './components/VoteBox';

function App() {
  const CANDIDATE_URL = import.meta.env.VITE_CANDIDATE_URL
  const ELECTION_URL = import.meta.env.VITE_ELECTION_URL

  const [candidates, setCandidates] = useState([])
  const [candidate, setCandidate] = useState({})
  const [electionStatus, setElectionStatus] = useState({})
  const [showVoteBox, setShowVoteBox] = useState(false)

  const fetchElectionStatus = async () => {
    const response = await axios(ELECTION_URL+"/status");
    setElectionStatus(response.data);
  };

  const fetchCandidates = async () => {
    const response = await axios(CANDIDATE_URL);
    setCandidates(response.data);
  };

  useEffect(() => {
    fetchCandidates();
    fetchElectionStatus();
  }, []);

  function voteBoxHandler(candidate) {
    setShowVoteBox(!showVoteBox)
    setCandidate(candidate)
    fetchElectionStatus();
    fetchCandidates();
  }

  return (
    <div>
      {showVoteBox && <VoteBox candidate={candidate} closeVoteBox={() => voteBoxHandler()} />}

      {!showVoteBox && <div className="container mx-auto text-center py-8">
        <h1 className='text-6xl mb-8 text-blue-700 font-bold'>Election Application</h1>

        <div className='grid gap-8 grid-cols-1 lg:grid-cols-3 xl:grid-cols-4'>
          {candidates.map((candidate) => (
            <Candidate 
              candidate={candidate} 
              key={candidate.ID} 
              electionStatus={electionStatus.enable}
              openVoteBox={() => voteBoxHandler(candidate)}
            />
            ))}
        </div>
      </div>}
    </div>
  )
}

export default App
