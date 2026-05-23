import { useNavigate } from '@umijs/max';
import { useEffect } from 'react';
import { useAuth } from 'react-oidc-context';

export default function Login() {
  const auth = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (auth.isAuthenticated) {
      navigate('/');
    } else if (!auth.isLoading) {
      auth.signinRedirect();
    }
  }, [auth, navigate]);

  if (auth.isLoading) {
    return (
      <div style={{ padding: '20px', textAlign: 'center' }}>加载中...</div>
    );
  }

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '80vh' }}>
      正在跳转到登录页面......
    </div>
  );
}