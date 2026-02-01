import { useState } from 'react'
import './App.css'
import { Tabs, type Tab } from './components/tabs'

function App() {
  const [tab,setTab] = useState<Tab>("Leaderboard")

  return <div className='h-screen bg-background'>
    <Tabs value={tab} onChange={setTab}/>
  </div>
}

export default App
