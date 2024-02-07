import React from 'react';
import ReactDOM from 'react-dom/client';
import { Home } from './Home.tsx';
import { Login } from './Login.tsx';
import './index.css';

import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { loader as userGoalsLoader, UserGoals } from './UserGoals.tsx';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Home />,
  },
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/:userId/goals',
    element: <UserGoals />,

    loader: userGoalsLoader,
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
