import './globals.css'
import { Inter } from 'next/font/google'
import Nav from '@/components/nav'
import AuthStatus from "../components/authStatus"
import SessionProviderWrapper from '@/utils/sessionProviderWrapper'

const inter = Inter({ subsets: ['latin'] })

export const metadata = {
  title: 'My demo',
  description: 'Some description for my website',
}

export default function RootLayout({ children }) {
  return (
    <SessionProviderWrapper>
    <html lang="en">
      <body className={inter.className}>
        <div className="flex flex-row">
          <div className="w-4/5 p-3 h-screen bg-black">{children}</div>
          <div className="w-1/5 p-3 h-screen bg-gray-700">
            <h2 className="text-3xl">Demo - frontend</h2>
              <AuthStatus />
            <hr />
              <Nav />
          </div>
        </div>
      </body>
    </html>
    </SessionProviderWrapper>
  )
}
