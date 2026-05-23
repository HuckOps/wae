import { useAuth } from 'react-oidc-context';

import { history } from '@umijs/max';
import BaseLayout from './base';
export default function () {
  const auth = useAuth();

  if (!auth || auth.isLoading) {
    return (
      <div>
        <div>加载中...</div>
      </div>
    );
  }

  if (!auth.isAuthenticated) {
    history.push('/login');
  }

  return (
    <div>
      <BaseLayout />
    </div>
  );
}
