import React, {useState} from "react";
import BaseForm, {BaseFormProps} from "./BaseForm";
import {useAuth} from "components/services/auth/AuthProvied";
import {useNavigate} from "react-router-dom";
import {toast, ToastContainer} from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export default function Register() {
    const {register} = useAuth();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState<string | null>(null);
    const [success, setSuccess] = useState<string | null>(null);
    const navigate = useNavigate();

    const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    };

    const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value);
    };

    const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setConfirmPassword(e.target.value);
    };

    const handleSubmit = async () => {
        if (password !== confirmPassword) {
            console.error("Passwords don't match");
            toast.error('Пароли не совпадают', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
            setError('Пароли не совпадают');
            return;
        }

        try {
            await register(email, password, confirmPassword);
            setError(null);
            setSuccess('Регистрация прошла успешно!');

            toast.success('Регистрация прошла успешно! Перенаправляем на страницу входа...', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });

            setTimeout(() => navigate('/login'), 3000);
        } catch (err) {
            console.error('Registration failed', err);
            setError('Ошибка регистрации. Попробуйте снова.');
            setSuccess(null);

            toast.error('Ошибка регистрации. Попробуйте снова.', {
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

    const formData: BaseFormProps = {
        title: "Регистрация",
        inputs: [
            {
                required: true,
                type: "email",
                placeholder: "Введите email",
                value: email,
                onChange: handleEmailChange
            },
            {
                required: true,
                type: "password",
                placeholder: "Введите пароль",
                value: password,
                onChange: handlePasswordChange
            },
            {
                required: true,
                type: "password",
                placeholder: "Подтвердите пароль",
                value: confirmPassword,
                onChange: handleConfirmPasswordChange
            },
        ],
        button: {
            text: "Зарегистрироваться",
        },
    };

    return (
        <>
            <BaseForm props={formData} onSubmit={handleSubmit}/>
            <ToastContainer/>
        </>
    );
}
