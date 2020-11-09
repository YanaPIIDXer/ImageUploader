import React from 'react';
import logo from './logo.svg';
import './App.css';
import UploadButton from './components/upload_button';

function App() {
  return (
    <div className="App">
      <h1>Upload your Image</h1>
      <h2>File should be Jpeg.</h2>

      <UploadButton />
    </div>
  );
}

export default App;
