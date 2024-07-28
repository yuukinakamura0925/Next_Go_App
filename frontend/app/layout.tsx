import './globals.css';
import { ReactNode } from 'react';

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <head>
        <title>My App</title>
      </head>
      <body>
        <nav className="bg-blue-600 p-4 text-white">
          <ul className="flex space-x-4">
            <li><a href="/auth/sign_up">新規登録</a></li>
            <li><a href="/auth/sign_in">サインイン</a></li>
            <li><a href="/books">本追加</a></li>
          </ul>
        </nav>
        <main className="p-4">{children}</main>
      </body>
    </html>
  );
}
