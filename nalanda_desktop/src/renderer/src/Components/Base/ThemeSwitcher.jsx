import React, { useState } from 'react';
import useDarkTheme from '../../Hooks/SetTheme';
import { DarkModeSwitch } from 'react-toggle-dark-mode';

export default function ThemeSwitcher() {
  const [colorTheme, setTheme] = useDarkTheme();
  const [darkSide, setDarkSide] = useState(colorTheme === 'light' ? true : false);

  const toggleDarkMode = checked => {
    setTheme(colorTheme);
    setDarkSide(checked);
  };

  return (
    <DarkModeSwitch checked={darkSide} onChange={toggleDarkMode} sunColor='black' moonColor='black'/>
  );
}