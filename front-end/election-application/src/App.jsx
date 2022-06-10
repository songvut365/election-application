import Candidate from './components/Candidate';

function App() {

  return (
    <div className="container mx-auto text-center py-8">
      <h1 className='text-6xl mb-4'>Election Application</h1>

      <div className='grid gap-6 grid-cols-1
        lg:grid-cols-4'
      >
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
      </div>
    </div>
  )
}

export default App
