import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';

// 1. Import MantineProvider dan CSS-nya
import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css'; 
import '@mantine/hooks/styles.css';


const rootElement = document.getElementById('root');
if (!rootElement) throw new Error('Failed to find the root element');

const root = createRoot(rootElement);

root.render(
  <StrictMode>
    <MantineProvider>
      <App />
    </MantineProvider>
  </StrictMode>
);