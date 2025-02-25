import FormLoginAdmin from "./components/FormLoginAdmin";

export default function FeatureAdminLogin() {
  return (
    <div className="w-10/12 rounded-lg py-6 shadow-lg">
      <p className="pb-2 font-semibold text-xl text-center">
        Admin Login
      </p>
      <FormLoginAdmin />
    </div>
  )
}
