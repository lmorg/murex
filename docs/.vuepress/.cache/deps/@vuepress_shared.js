import {
  isArray,
  isFunction,
  isString
} from "./chunk-L52MBEYQ.js";

// node_modules/.pnpm/@vuepress+shared@2.0.0-beta.66/node_modules/@vuepress/shared/dist/index.js
var resolveHeadIdentifier = ([
  tag,
  attrs,
  content
]) => {
  if (tag === "meta" && attrs.name) {
    return `${tag}.${attrs.name}`;
  }
  if (["title", "base"].includes(tag)) {
    return tag;
  }
  if (tag === "template" && attrs.id) {
    return `${tag}.${attrs.id}`;
  }
  return JSON.stringify([tag, attrs, content]);
};
var dedupeHead = (head) => {
  const identifierSet = /* @__PURE__ */ new Set();
  const result = [];
  head.forEach((item) => {
    const identifier = resolveHeadIdentifier(item);
    if (!identifierSet.has(identifier)) {
      identifierSet.add(identifier);
      result.push(item);
    }
  });
  return result;
};
var ensureLeadingSlash = (str) => str[0] === "/" ? str : `/${str}`;
var ensureEndingSlash = (str) => str[str.length - 1] === "/" || str.endsWith(".html") ? str : `${str}/`;
var formatDateString = (str, defaultDateString = "") => {
  const dateMatch = str.match(/\b(\d{4})-(\d{1,2})-(\d{1,2})\b/);
  if (dateMatch === null) {
    return defaultDateString;
  }
  const [, yearStr, monthStr, dayStr] = dateMatch;
  return [yearStr, monthStr.padStart(2, "0"), dayStr.padStart(2, "0")].join("-");
};
var isLinkFtp = (link) => link.startsWith("ftp://");
var isLinkHttp = (link) => /^(https?:)?\/\//.test(link);
var markdownLinkRegexp = /.md((\?|#).*)?$/;
var isLinkExternal = (link, base = "/") => {
  if (isLinkHttp(link) || isLinkFtp(link)) {
    return true;
  }
  if (link.startsWith("/") && !link.startsWith(base) && !markdownLinkRegexp.test(link)) {
    return true;
  }
  return false;
};
var isLinkMailto = (link) => /^mailto:/.test(link);
var isLinkTel = (link) => /^tel:/.test(link);
var isPlainObject = (val) => Object.prototype.toString.call(val) === "[object Object]";
var omit = (obj, ...keys) => {
  const result = { ...obj };
  for (const key of keys) {
    delete result[key];
  }
  return result;
};
var removeEndingSlash = (str) => str[str.length - 1] === "/" ? str.slice(0, -1) : str;
var removeLeadingSlash = (str) => str[0] === "/" ? str.slice(1) : str;
var resolveLocalePath = (locales, routePath) => {
  const localePaths = Object.keys(locales).sort((a, b) => {
    const levelDelta = b.split("/").length - a.split("/").length;
    if (levelDelta !== 0) {
      return levelDelta;
    }
    return b.length - a.length;
  });
  for (const localePath of localePaths) {
    if (routePath.startsWith(localePath)) {
      return localePath;
    }
  }
  return "/";
};
var resolveRoutePathFromUrl = (url, base = "/") => {
  const pathname = url.replace(/^(https?:)?\/\/[^/]*/, "");
  return pathname.startsWith(base) ? `/${pathname.slice(base.length)}` : pathname;
};
export {
  dedupeHead,
  ensureEndingSlash,
  ensureLeadingSlash,
  formatDateString,
  isArray,
  isFunction,
  isLinkExternal,
  isLinkFtp,
  isLinkHttp,
  isLinkMailto,
  isLinkTel,
  isPlainObject,
  isString,
  omit,
  removeEndingSlash,
  removeLeadingSlash,
  resolveHeadIdentifier,
  resolveLocalePath,
  resolveRoutePathFromUrl
};
//# sourceMappingURL=@vuepress_shared.js.map
