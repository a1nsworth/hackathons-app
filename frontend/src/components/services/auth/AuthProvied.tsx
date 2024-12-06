import React, {createContext, useState, useContext, useEffect, ReactNode} from 'react';
import axios from 'axios';
import {AuthUrl} from "config";

// Типы для данных пользователя и токенов
interface User {
    id: string;
    email: string;
}

interface AuthContextType {
    user: User | null;
    isAuthenticated: boolean;
    accessToken: string | null;
    refreshToken: string | null;
    login: (email: string, password: string) => Promise<void>;
    register: (email: string, password: string, confirmPassword: string) => Promise<void>;
    logout: () => void;
    refreshAuthToken: () => Promise<void>;
}

// Контекст для хранения данных авторизации
const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = (): AuthContextType => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};

export const AuthProvider = ({children}: { children: ReactNode }) => {
    const [user, setUser] = useState<User | null>(null);
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [accessToken, setAccessToken] = useState<string | null>(localStorage.getItem('accessToken') || null);
    const [refreshToken, setRefreshToken] = useState<string | null>(localStorage.getItem('refreshToken') || null);

    const login = async (email: string, password: string) => {
        try {
            console.log("make response")
            const response = await axios.post(AuthUrl.Login, {email, password});
            console.log(response.data.accessToken);
            console.log(response.data.refreshToken);
            setAccessToken(response.data.accessToken);
            setRefreshToken(response.data.refreshToken);
            localStorage.setItem('accessToken', response.data.accessToken);
            localStorage.setItem('refreshToken', response.data.refreshToken);
            setIsAuthenticated(true)
            // setUser(response.data.user);
        } catch (error) {
            throw error;
        }
    };
    const register = async (email: string, password: string, confirmPassword: string) => {
        try {
            const response = await axios.post(AuthUrl.Register, {email, password, confirmPassword});
            // После успешной регистрации сразу выполняем логин
            setAccessToken(response.data.accessToken);
            setRefreshToken(response.data.refreshToken);
            setIsAuthenticated(true)
            localStorage.setItem('accessToken', response.data.accessToken);
            localStorage.setItem('refreshToken', response.data.refreshToken);
            setUser(response.data.user);
        } catch (error) {
            console.error('Registration failed', error);
            setIsAuthenticated(false)
            throw error;
        }
    };

    const refreshAuthToken = async () => {
        if (!accessToken) {
            console.error('No access token available');
        }
        if (!refreshToken) {
            console.error('No refresh token available');
            return;
        }

        try {
            const response = await axios.post(AuthUrl.Refresh, {accessToken, refreshToken});
            console.log(response.data.accessToken);
            setIsAuthenticated(true)
            setAccessToken(response.data.accessToken);
            localStorage.setItem('accessToken', response.data.accessToken);
        } catch (error) {
            console.error('Refresh token failed', error);
            logout();
            throw error;
        }
    };

    const logout = () => {
        setAccessToken(null);
        setRefreshToken(null);
        setUser(null);
        setIsAuthenticated(false);
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
    };

    useEffect(() => {
        if (accessToken) {
            axios.defaults.headers['Authorization'] = `Bearer ${accessToken}`;
        }
    }, [accessToken]);

    return (
        <AuthContext.Provider
            value={{user, isAuthenticated, accessToken, refreshToken, login, register, logout, refreshAuthToken}}>
            {children}
        </AuthContext.Provider>
    );
};
