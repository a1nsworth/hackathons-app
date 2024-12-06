import {Navigate, useLocation} from "react-router-dom";
import {useAuth} from "./AuthProvied";

interface PrivateRouteProps {
    children: JSX.Element;
    redirectTo?: string;
    condition?: () => boolean;
}

export default function PrivateRoute({
                                         children,
                                         redirectTo = "/login",
                                         condition,
                                     }: PrivateRouteProps): JSX.Element {
    const {isAuthenticated} = useAuth();
    const location = useLocation();

    console.log("isAuth:", isAuthenticated); // Проверьте, есть ли пользователь
    console.log("Condition:", condition ? condition() : "No condition");

    if (!isAuthenticated || (condition && !condition())) {
        return <Navigate to={redirectTo} state={{from: location}}/>;
    }

    return children;
}
