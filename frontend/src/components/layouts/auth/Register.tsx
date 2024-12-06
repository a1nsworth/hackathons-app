import BaseForm, {BaseFormProps} from "./BaseForm";
import {useState} from "react";
import {useAuth} from "components/services/auth/AuthProvied";

export default function Register() {
    const {register} = useAuth()
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState<string | null>(null);

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
            setError('Passwords do not match');
            return;
        }

        try {
            console.log('Registration data:', {email, password, confirmPassword});
            await register(email, password, confirmPassword)
            setError(null); // сбросить ошибку
        } catch (err) {
            console.error('Registration failed', err);
        }
    };

    const formData: BaseFormProps = {
        title: "Register",
        inputs: [
            {required: true, type: "email", placeholder: "Input email", value: email, onChange: handleEmailChange},
            {
                required: true,
                id: "password",
                type: "password",
                placeholder: "Input password",
                value: password,
                onChange: handlePasswordChange
            },
            {
                required: true,
                id: "conf-password",
                type: "password",
                placeholder: "Confirm password",
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
            {error && <p style={{color: 'red'}}>{error}</p>}
            <BaseForm props={formData} onSubmit={handleSubmit}/>
        </>
    );
}