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

export default function Candidate() {
  let candidate =   {
    id: "1",
    name: "Elon Musk",
    dob: "June 28, 1971",
    bioLink: "https://en.wikipedia.org/wiki/Elon_Musk",
    imageLink: "https://upload.wikimedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg",
    policy: "Choose me if your don't know who to choose",
    votedCount: 1100
  }

  return (
    <div className='max-w-sm rounded-lg shadow-lg p-4'>
      <div className='relative'>
        <img src={candidate.imageLink} className="w-full h-96 object-cover object-top rounded-lg shadow-md" />
        <h1 className="absolute top-0 right-0 px-4 py-2 bg-gray-200 text-5xl font-bold rounded-tr-lg">
          {candidate.id}
        </h1>
      </div>

      <div className='flex justify-between p-4'>
        <div>
          <h3 className='font-bold text-lg'>{candidate.name}</h3>
          <p className='text-xs font-bold text-gray-400'>{ageFromDate(candidate.dob)} yrs</p>
        </div>

        <div>
          <h2 className='font-bold text-xl'>{formatVotedCount(candidate.votedCount)}</h2>
          <p className='text-sm'>votes</p>
        </div>
      </div>
      
      <h2 className='text-2xl p-4'>" {candidate.policy} "</h2>

      <Button>VOTE</Button>
    </div>
  )
}
