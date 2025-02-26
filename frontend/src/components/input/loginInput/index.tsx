import { TextField } from "@mui/material";
import React from "react";

interface ComponentProps {
  name: string
  label: string
  size: "small" | "medium"
  placeholder: string
  type: string
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const InputLogin: React.FC<ComponentProps> = ({ name, label, size, placeholder, type, onChange }) => {
  return (
    <
      TextField
      className="w-4/6"
      name={name}
      label={label}
      size={size}
      id="outlined-size-normal"
      placeholder={placeholder}
      type={type ? type : "text"}
      onChange={onChange}
    />
  )
}

export default InputLogin
