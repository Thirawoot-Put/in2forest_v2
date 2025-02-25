import { TextField } from "@mui/material";

type TProps = {
  label: string
  size: "small" | "medium"
  placeholder: string
  type: string
}

export default function InputLogin({ label, size, placeholder, type }: TProps) {
  return (
    <
      TextField
      className="w-4/6"
      label={label}
      size={size}
      id="outlined-size-normal"
      placeholder={placeholder}
      type={type ? type : "text"}
    />
  )
}
