import FormLoginAdmin from "./components/FormLoginAdmin";

export default function FeatureAdminLogin() {
  return (
    <div className="w-10/12 border border-gray-500 rounded-lg py-6">
      <p className="pb-4 font-semibold text-xl text-center">Admin Login</p>
      <FormLoginAdmin />
    </div>
  )
}
