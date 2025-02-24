"use-client"

import { TextField } from "@mui/material";

export default function FormLoginAdmin() {
  return (
    <form>
      <div id="TextField1">
        <TextField label="Username" id="outlined-size-normal" placeholder="sample@mail.com" />
      </div>
      <div id="TextField2">
        <TextField label="Password" id="outlined-size-normal" type="password" placeholder="password1234" />
      </div>
    </form>
  )
}
