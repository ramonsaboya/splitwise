import React from 'react';
import { createUseStyles } from 'react-jss';

const useStyles = createUseStyles({
  container: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    height: '100vh',
    width: '100vw',
    margin: 0,
    padding: 0,
  },
});

export default function App() {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      Melhor que o Splitwise
    </div>
  );
}