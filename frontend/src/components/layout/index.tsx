import React from "react";

export default function LayoutComponent({ children }: { children: React.ReactNode }) {
  return (
    <main className="h-screen w-screen p-5" >
      {children}
    </main>
  )
}
