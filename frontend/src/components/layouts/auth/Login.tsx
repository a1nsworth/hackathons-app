import React, {useState} from "react";
import BaseForm, {BaseFormProps} from "./BaseForm";
import {useAuth} from "components/services/auth/AuthProvied";
import {toast, ToastContainer} from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export default function Login() {
    const {login} = useAuth();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    };

    const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value);
    };

    const handleSubmit = async () => {
        try {
            console.log("Logging in with", {email, password});
            await login(email, password);
            // Уведомление об успешном входе
            toast.success("Успешный вход!", {
                position: "top-right",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
        } catch (error) {
            console.error("Login failed", error);
            // Уведомление об ошибке
            toast.error("Ошибка входа. Проверьте данные и попробуйте снова.", {
                position: "top-right",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
        }
    };

    const updatedLoginData: BaseFormProps = {
        title: "Login",
        inputs: [
            {
                required: true,
                type: "email",
                placeholder: "Input email",
                value: email,
                onChange: handleEmailChange,
            },
            {
                required: true,
                type: "password",
                placeholder: "Input password",
                value: password,
                onChange: handlePasswordChange,
            },
        ],
        button: {
            text: "Войти",
        },
    };

    return (
        <>
            <BaseForm props={updatedLoginData} onSubmit={handleSubmit}/>
            <ToastContainer/>
        </>
    );
}
