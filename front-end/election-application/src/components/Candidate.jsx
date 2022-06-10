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
    <div className='max-w-sm rounded-lg shadow-lg p-4 bg-white relative'>
      <div className='relative'>
        <img src={candidate.imageLink} className="w-full h-96 object-cover object-top rounded-lg shadow-md" />
        <h1 className="absolute top-0 right-0 px-4 py-2 bg-gray-200 text-5xl font-bold rounded-tr-md">
          {candidate.id}
        </h1>
      </div>

      <div className='flex justify-between p-4'>
        <div className='text-left'>
          <h3 className='font-bold text-lg'>{candidate.name}</h3>
          <p className='text-xs font-bold text-gray-400'>{ageFromDate(candidate.dob)} yrs</p>
        </div>

        <div>
          <h2 className='font-bold text-xl'>{formatVotedCount(candidate.votedCount)}</h2>
          <p className='text-sm'>votes</p>
        </div>
      </div>
      
      <h2 className='text-2xl p-4 mb-4'>" {candidate.policy} "</h2>

      <Button className="absolute bottom-0">VOTE</Button>
    </div>
  )
}
