import Candidate from './components/Candidate';

function App() {
  let candidates =  [
    {
      id: "1",
      name: "Elon Musk",
      dob: "June 28, 1971",
      bioLink: "https://en.wikipedia.org/wiki/Elon_Musk",
      imageLink: "https://upload.wikimedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg",
      policy: "Choose me if your don't know who to choose",
      votedCount: 1100
    },
    {
      id: "2",
      name: "Songvut Nakrong",
      dob: "June 28, 1998",
      bioLink: "https://en.wikipedia.org/wiki/Elon_Musk",
      imageLink: "https://scontent.fbkk5-4.fna.fbcdn.net/v/t1.6435-9/106128727_2683554015300421_102285652757321707_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=09cbfe&_nc_eui2=AeEoeyjaf9l7h2UubPHG1q7qUQdlVNY2GitRB2VU1jYaK76b-yGJRvyPjwM4LRagURu_83_imCAsdynhN-AFolJm&_nc_ohc=hNcItozul_4AX9hs-WR&_nc_ht=scontent.fbkk5-4.fna&oh=00_AT_MYaF-d0N3d4TFXkl0ADdCiICV36P2kqVuJKB6m6jIfQ&oe=62C9C4DE",
      policy: "I don't know why I am here but I am",
      votedCount: 232
    },
    {
      "id": "3",
      "name": "Jeff Bezos",
      "dob": "January 12, 1964",
      "bioLink": "https://en.wikipedia.org/wiki/Jeff_Bezos",
      "imageLink": "https://pbs.twimg.com/profile_images/669103856106668033/UF3cgUk4_400x400.jpg",
      "policy":"Choose me if your don't know who to choose",
      "votedCount": 0
    },
    {
      "id": "4",
      "name": "Brown",
      "dob": "August 8, 2011",
      "bioLink": "https://line.fandom.com/wiki/Brown",
      "imageLink": "https://cdn.shopify.com/s/files/1/0231/6137/2752/files/bf_brown.jpg?v=1622860193",
      "policy": "Brown, Brown, Brown , Brown and Brown"
    }
  ]

  return (
    <div className="container mx-auto text-center py-8">
      <h1 className='text-6xl mb-8 text-blue-700 font-bold'>Election Application</h1>

      <div className='grid gap-8 grid-cols-1 lg:grid-cols-4'>
        {candidates.map(candidate => (
          <Candidate candidate={candidate} />
        ))}
      </div>
    </div>
  )
}

export default App
