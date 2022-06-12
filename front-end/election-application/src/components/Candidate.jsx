import React from 'react';

import Button from "./Button";

function ageFromDate(birthDate) {
  let birthYear = birthDate.split(",")[1]
  let date = new Date()
  let thisYear = date.getFullYear()
  let age = thisYear - birthYear
  return age
}

function formatVotedCount(votedCount) {
  if (votedCount >= 1000 && votedCount < 1000000) {
    return votedCount / 1000 + "k"
  }
  else if (votedCount >= 1000000) {
    return votedCount / 1000000 + "M"
  }
  return votedCount
}

export default function Candidate(props) {
  let candidate = props.candidate;
  return (
    <div className='max-w-sm rounded-lg shadow-lg mx-4 p-4 bg-white relative flex flex-col justify-between'>
      <div className='relative'>
        <div className="hover:scale-105">
          <img src={candidate.ImageLink} 
            className="w-full h-96 object-cover object-top rounded-lg shadow-md" 
            />
          <h1 className="absolute top-0 right-0 px-4 py-2 bg-gray-200 text-5xl font-bold rounded-tr-md">
            {candidate.ID}
          </h1>
        </div>

        <div className='flex justify-between p-4'>
          <div className='text-left'>
            <h3 className='font-bold text-lg'>{candidate.Name}</h3>
            <p className='text-xs font-bold text-gray-400'>{ageFromDate(candidate.DOB)} yrs</p>
          </div>

          <div>
            <h2 className='font-bold text-xl'>{formatVotedCount(candidate.VotedCount)}</h2>
            <p className='text-sm'>votes</p>
          </div>
        </div>

        <h2 className='text-2xl p-4 mb-4'>" {candidate.Policy} "</h2>
      </div>

      {props.electionStatus ? (
        <Button 
          method={props.openVoteBox}
          bgColor="blue" textColor="white"
        >  
          VOTE
        </Button>
      ) : (
        <div className="my-4 w-full bg-gray-200 h-8 relative rounded-sm">
          <div className="bg-blue-600 h-8 absolute z-10 rounded-sm border-2 box-border" style={{ width: `${candidate.Percentage}`}}>
            <p className='w-full h-8 absolute z-20 text-center pl-2 pt-0.5 font-bold text-white '>
                {candidate.Percentage}
            </p>
          </div>
        </div>
      )}

      
    </div>
  )
}
