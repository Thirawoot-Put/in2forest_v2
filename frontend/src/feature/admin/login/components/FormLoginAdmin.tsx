"use-client"

import { TextField } from "@mui/material";

export default function FormLoginAdmin() {
  return (
    <form className="flex flex-col gap-2 items-center w-full h-full py-4 ">
      <TextField className="w-4/6" label="Username" size="small" id="outlined-size-normal" placeholder="sample@mail.com" />
      <TextField className="w-4/6" label="Password" size="small" id="outlined-size-normal" type="password" placeholder="password1234" />
    </form>
  )
}
