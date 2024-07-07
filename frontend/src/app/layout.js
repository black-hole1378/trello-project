// app/layout.js
"use client";

import { usePathname } from "next/navigation";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "../component/header/header";
import Dashboard from "../component/dashboard/dashboard";
import { Provider } from "react-redux";
import store from "../store";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({ children }) {
  const pathname = usePathname();
  const showHeader = pathname !== "/login" && pathname !== "/signup";

  return (
    <html lang="en">
      <body className={inter.className}>
        <Provider store={store}>
          {showHeader && <Header />}
          {children}
        </Provider>
      </body>
    </html>
  );
}
