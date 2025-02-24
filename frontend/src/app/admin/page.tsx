import LayoutComponent from "@/components/layout";
import FeatureAdminLogin from "@/feature/admin/login";

export default function AdminPage() {
  return (
    <LayoutComponent>
      <div className="w-full h-full flex justify-center items-center">
        <FeatureAdminLogin />
      </div>
    </LayoutComponent>
  )
}
