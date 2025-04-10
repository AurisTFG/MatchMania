import React from "react";
import "./Content.css";

interface ContentProps {
  children: React.ReactNode;
}

export default function Content({ children }: ContentProps) {
  return <div className="content-container">{children}</div>;
}
