import React from 'react';
import ReactDOM from 'react-dom/client';
import { Home } from './Home.tsx';
import { Login } from './Login.tsx';
import './index.css';

import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { loader as userGoalsLoader, UserGoals } from './UserGoals.tsx';
import { UserGoalsNew } from './UserGoalsNew.tsx';
import { Register } from './Register.tsx';
import { AuthProvider } from './auth/AuthProvider.tsx';
import {
  loader as userActiviesNewLoader,
  UserActivitiesNew,
} from './UserActivitiesNew.tsx';

const router = createBrowserRouter([
  {
    element: <AuthProvider />,
    children: [
      {
        path: '/',
        element: <Home />,
      },
      {
        path: '/register',
        element: <Register />,
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
      {
        path: '/:userId/goals/new',
        element: <UserGoalsNew />,
      },
      {
        path: '/:userId/activities/new',
        element: <UserActivitiesNew />,

        loader: userActiviesNewLoader,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
