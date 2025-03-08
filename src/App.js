import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Home from './pages/Home';
import Detail from './pages/Detail';
import Form from './pages/Form';
import Logs from './pages/Logs';
import GlobalButtons from './components/GlobalButtons';
import Login from './pages/Login';
import PrivateRoute from './components/PrivateRoute';
import { AuthProvider } from './components/AuthContext'; 
import './styles/styles.css';  // стили

/* <button id="staffers">Список сотрудников</button>
        <button id="logs">Логи проходов</button>
      </GlobalButtons> 
      
          <Route path="/path1" element={<Stuffers />} />*/
const App = () => {
  return (
    <AuthProvider>
      <Router>
        <GlobalButtons />
        <Routes>
        <Route path="/login" element={<Login />} />
          <Route path="/" element={<PrivateRoute><Home /></PrivateRoute>} />
          <Route path="/detail/:id" element={<PrivateRoute><Detail /></PrivateRoute>} />
          <Route path="/add" element={<PrivateRoute><Form /></PrivateRoute>} />
          <Route path="/logs" element={<PrivateRoute><Logs /></PrivateRoute>} />
        </Routes>
      </Router>
    </AuthProvider>
  );
};

export default App;
