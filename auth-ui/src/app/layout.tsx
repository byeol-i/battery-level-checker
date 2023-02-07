import './globals.css'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html>
      <head /> 
        {/* head.tsx. Find out more at https://beta.nextjs.org/docs/api-reference/file-conventions/head */}
      <head />
      <body>{children}</body>
    </html>
  )
}
