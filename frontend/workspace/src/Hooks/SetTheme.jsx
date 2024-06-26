import { useState, useEffect } from "react";

export default function useDarkTheme() {
    const [theme, setTheme] = useState(localStorage.theme || "light");
    const colorTheme = theme === "dark" ? "light": "dark";
    useEffect(() => {
      const root = window.document.documentElement;
      root.classList.remove(colorTheme);
      root.classList.add(theme);
      // Use local storage.
      localStorage.setItem('theme', theme);
    }, [theme, colorTheme]);

    return [colorTheme,setTheme];
}