import React, {useState, useEffect} from 'react';
import axios from 'axios';

import Candidate from '../components/Candidate';
import VoteBox from '../components/VoteBox';

export default function Home() {
  const ELECTION_URL = import.meta.env.VITE_ELECTION_URL

  const [candidates, setCandidates] = useState([])
  const [candidate, setCandidate] = useState({})
  const [mayor, setMayor] = useState({})
  const [electionStatus, setElectionStatus] = useState({})
  const [showVoteBox, setShowVoteBox] = useState(false)

  const fetchElectionStatus = async () => {
    await axios(ELECTION_URL+"/status").then(response => {
      setElectionStatus(response.data);

      if (response.data.enable === false) {
        findNewMayor();
      }
    });
  };

  const fetchCandidates = async () => {
    await axios.post(ELECTION_URL+"/result").then(response => {
      setCandidates(response.data);
    })
  };

  function voteBoxHandler(candidate) {
    setShowVoteBox(!showVoteBox)
    setCandidate(candidate)
    fetchCandidates();
    fetchElectionStatus();
  }

  function findNewMayor() {
    axios.post(ELECTION_URL+"/result").then(response => {
      let allCandidates = response.data
      let max =  0
      for (let c of allCandidates) {
        let percentageInt = parseInt(c.Percentage.substring(0,3))
        if (max < percentageInt) {
          max = percentageInt
          setMayor(c)
        }
      }
    })
  }

  useEffect(() => {
    fetchCandidates();
    fetchElectionStatus();
  }, []);

  return (
    <div>
    {showVoteBox && <VoteBox candidate={candidate} closeVoteBox={() => voteBoxHandler()} />}

    {!showVoteBox && <div className="container mx-auto text-center py-8">
      <h1 className='text-6xl mb-8 text-blue-700 font-bold'>Election Application</h1>

      {!electionStatus.enable && <p className='mb-8 text-xl'>
        The new mayor is: <span className="text-3xl font-bold italic">
          #{mayor.ID} {mayor.Name}
        </span>
      </p>}

      <div className='grid gap-8 grid-cols-1 lg:grid-cols-3 xl:grid-cols-4'>
        {candidates.map((candidate) => (
          <Candidate 
            candidate={candidate} 
            key={candidate.Name} 
            electionStatus={electionStatus.enable}
            openVoteBox={() => voteBoxHandler(candidate)}
          />
          ))}
      </div>
    </div>}
  </div>
  )
}
