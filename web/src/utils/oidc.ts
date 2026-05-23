import { User } from "oidc-client-ts";
import { baseRequest } from "./baseRequest";

interface OIDCConfig {
  provider: string;
  client_id: string;
}

const CACHE_KEY = "oidc_config";

let cachedPromise: Promise<OIDCConfig> | null = null;

function saveToStorage(config: OIDCConfig) {
  localStorage.setItem(CACHE_KEY, JSON.stringify(config));
}

function loadFromStorage(): OIDCConfig | null {
  const cached = localStorage.getItem(CACHE_KEY);
  if (cached) {
    try {
      return JSON.parse(cached) as OIDCConfig;
    } catch {
      return null;
    }
  }
  return null;
}

export async function getOIDCConfig(): Promise<OIDCConfig> {
  if (cachedPromise) {
    return cachedPromise;
  }

  const cached = loadFromStorage();
  if (cached) {
    cachedPromise = Promise.resolve(cached);
    return cachedPromise;
  }

  cachedPromise = baseRequest
    .get("/v1/oidc/config")
    .then((result: any) => {
      if (result.code !== 0) {
        throw new Error(result.message || "Failed to fetch OIDC config");
      }
      const config = result.data as OIDCConfig;
      saveToStorage(config);
      return config;
    })
    .catch((error) => {
      cachedPromise = null;
      throw error;
    });

  return cachedPromise;
}

export async function getUser(): Promise<User | null> {
  const config = await getOIDCConfig();
  const oidcStorage = localStorage.getItem(
    `oidc.user:${config.provider}:${config.client_id}`
  );
  if (!oidcStorage) {
    return null;
  }

  return User.fromStorageString(oidcStorage);
}

export async function removeUser(): Promise<void> {
  const config = await getOIDCConfig();
  localStorage.removeItem(
    `oidc.user:${config.provider}:${config.client_id}`
  );
}