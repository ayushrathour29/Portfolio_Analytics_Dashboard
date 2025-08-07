import React from 'react';
import Dashboard from './pages/Dashboard';
import './App.css'

function App() {


  return (
    <>
       <div className="min-h-screen bg-gray-50 text-gray-800">
      <header className="p-4 bg-blue-600 text-white text-2xl font-bold text-center">
        Portfolio Analytics Dashboard
      </header>
      <main className="max-w-6xl mx-auto py-8">
        <Dashboard />
      </main>
    </div>
     
    </>
  )
}

export default App
