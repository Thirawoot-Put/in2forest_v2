"use-client"

import InputLogin from "@/components/input/loginInput";
import { Button, Divider } from "@mui/material";

export default function FormLoginAdmin() {
  return (
    <form className="w-full h-full py-4">
      <div className="flex flex-col gap-2 items-center ">
        <InputLogin label="Username" size="small" placeholder="sample@mail.com" type="" />
        <InputLogin label="Password" size="small" placeholder="password1234" type="password" />
      </div>
      <div className="flex flex-col items-center justify-center gap-2 pt-5">
        <Button variant="contained">Login</Button>
        <Divider className="w-3/5">or</Divider>
        <Button variant="text">Register</Button>
      </div>
    </form >
  )
}
