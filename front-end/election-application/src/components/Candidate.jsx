import React from 'react';
import './Candidate.css';

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
    <div className='candidate'>
      <div className='imageBox'>
        <img src={candidate.imageLink} className="candidateImage" />
        <h1 className="candidateId">{candidate.id}</h1>
      </div>

      <div className='detailBox'>
        <div className='infoBox'>
          <h3 className='candidateName'>{candidate.name}</h3>
          <p className='candidateAge'>{ageFromDate(candidate.dob)} yrs</p>
        </div>

        <div className='voteBox'>
          <h2 className='candidateVotedCount'>{formatVotedCount(candidate.votedCount)}</h2>
          <p className='candidateVoteLabel'>votes</p>
        </div>
      </div>
      
      <h2 className='candidatePolicy'>" {candidate.policy} "</h2>

      <button className='voteButton'>VOTE</button>
    </div>
  )
}
