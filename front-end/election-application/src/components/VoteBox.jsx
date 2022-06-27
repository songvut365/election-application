import React, {useState} from 'react';
import axios from 'axios';

import Button from './Button';

export default function VoteBox(props) {
  const VOTE_URL = import.meta.env.VITE_VOTE_URL;
  const [nationalID, setNationalID] = useState("");
  const [votedSuccess, setVotedSuccess] = useState(false);
  const [voteStatus, setVoteStatus] = useState("not voted yet")

  const vote = async (candidate) => {
    const data = {
      "NationalID": nationalID,
      "CandidateID": candidate.ID
    }
    const response = await axios.post(VOTE_URL, data)
    if (response.data.status === "ok") {
      setVoteStatus("Thank you for your voted")
      setVotedSuccess(true)
    } 
    else if (response.data.message === "Already voted") {
      setVoteStatus("You have already voted")
      setVotedSuccess(true)
    }
  }

  return (
    <div className='absolute w-screen h-screen flex z-10'>
      <div className='max-w-sm h-2/3 rounded-lg shadow-lg p-6 bg-white m-auto'>
          {!votedSuccess ? (
            // Before vote
            <div className='flex relative flex-col justify-between h-full'>
              <div>
                <p className='text-3xl'>Please enter your national ID to confirm your vote</p>
                <input type="text" placeholder='x-xxxx-xxxxx-xx-x' maxLength={13}
                  value={nationalID} onChange={e => setNationalID(e.target.value)}
                  className='w-full h-12 border-2 my-4 py-2 px-2 rounded-md appearance-none
                    focus:outline-none focus:shadow-outline text-slate-500'/>
              </div>

              <div className='flex justify-between gap-4'>
                <Button bgColor="blue" textColor="white" method={() => vote(props.candidate)}>Vote</Button>
                <Button bgColor="white" method={props.closeVoteBox}>Cancel</Button>
              </div>
            </div>
          ) : (
            //After Vote
            <div className='flex relative flex-col justify-center h-full'>
              <p className='text-3xl pb-8'>{voteStatus}</p>
              <Button bgColor="blue" textColor="white" method={props.closeVoteBox}>Done</Button>
            </div>
          )}
        </div>
    </div>
  )
}