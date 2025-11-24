const DEFAULT_TIMEOUT = 15000;

function buildUrl(path) {
    const base = import.meta.env.VITE_API_BASE ?? '';
    if (/^https?:\/\//i.test(path)) return path;
    const _base = base.endsWith('/') ? base.slice(0, -1) : base;
    const _path = path.startsWith('/') ? path : `/${path}`;
    return `${_base}${_path}`;
}

export async function apiFetch(path, options = {}) {
    const url = buildUrl(path);
    const controller = new AbortController();
    const timeout = options.timeout ?? DEFAULT_TIMEOUT;
    const timer = setTimeout(() => controller.abort(), timeout);

    const fetchOptions = {
        headers: {
            'Content-Type': 'application/json',
            ...(options.headers || {}),
        },
        signal: controller.signal,
        ...options,
    };

    try {
        const res = await fetch(url, fetchOptions);
        clearTimeout(timer);
        if (!res.ok) {
            let body = null;
            try {
                body = await res.json();
            } catch (e) {
                body = await res.text();
            }
            const err = new Error(body && body.error ? body.error : `HTTP ${res.status}`);
            err.status = res.status;
            err.body = body;
            throw err;
        }

        const text = await res.text();
        try {
            return JSON.parse(text);
        } catch {
            return text;
        }
    } catch (err) {
        if (err.name === 'AbortError') {
            throw new Error('请求超时');
        }
        throw err;
    }
}

export default apiFetch;
