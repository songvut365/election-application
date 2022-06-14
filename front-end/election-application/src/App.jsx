import React, { useState } from 'react';
import { Routes, Route } from 'react-router-dom'

import Home from './views/Home';
import Stream from './views/Stream';
import Download from './views/Download';
import Login from './views/Login';
import Register from './views/Register';
import Profile from './views/Profile';

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="stream" element={<Stream />}/>
        <Route path="download" element={<Download />}/>
        <Route path="login" element={<Login />}/>
        <Route path="register" element={<Register />}/>
        <Route path="profile" element={<Profile />}/>
      </Routes>
    </div>
  )
}

export default App
