import { getOIDCConfig } from '@/utils/oidc';
import { WebStorageStateStore } from 'oidc-client-ts';
import { useEffect, useState } from 'react';
import { Toaster } from 'react-hot-toast';
import { AuthProvider } from 'react-oidc-context';
import Auth from './auth';

export default function Layout() {
  const [oidcConfig, setOidcConfig] = useState<any>(null);

  useEffect(() => {
    getOIDCConfig().then((config) => {
      setOidcConfig({
        authority: config.provider,
        client_id: config.client_id,
        redirect_uri: `${window.location.origin}/callback`,
        response_type: 'code',
        scope: 'openid profile email',
        automaticSilentRenew: true,
        loadUserInfo: true,
        userStore: new WebStorageStateStore({ store: window.localStorage }),
      });
    });
  }, []);

  if (!oidcConfig) {
    return <div>加载中...</div>;
  }

  return (
    <AuthProvider {...oidcConfig}>
      <Toaster
        position="top-center"
        toastOptions={{
          duration: 3000,
        }}
      />
      <Auth />
    </AuthProvider>
  );
}
