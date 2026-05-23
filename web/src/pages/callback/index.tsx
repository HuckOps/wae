import { useNavigate } from '@umijs/max';
import { useEffect } from 'react';
import { useAuth } from 'react-oidc-context';

export default function CallbackPage() {
  const auth = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (auth.isAuthenticated) {
      navigate('/');
    }
  }, [auth.isAuthenticated, navigate]);

  if (auth.error) {
    return (
      <div style={{ padding: '20px', textAlign: 'center' }}>
        登录失败: {auth.error.message}
      </div>
    );
  }

  return (
    <div style={{ padding: '20px', textAlign: 'center' }}>正在验证身份...</div>
  );
}
