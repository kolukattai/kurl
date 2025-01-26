
function markdownToHTML(markdown, dynamicClasses = {}) {
  console.log("markdown",markdown);
  const generateID = (text) => slugify(text);  // Create dynamic IDs
  
  // Convert headers and apply dynamic IDs and classes
  markdown = markdown.replace(/^###### (.*?)$/gm, (match, content) => {
      const id = generateID(content);
      return `<h6 id="${id}" class="${dynamicClasses.h6 || 'markdown-header'}">${content}</h6>`;
  });
  markdown = markdown.replace(/^##### (.*?)$/gm, (match, content) => {
      const id = generateID(content);
      return `<h5 id="${id}" class="${dynamicClasses.h5 || 'markdown-header'}">${content}</h5>`;
  });
  markdown = markdown.replace(/^#### (.*?)$/gm, (match, content) => {
      const id = generateID(content);
      return `<h4 id="${id}" class="${dynamicClasses.h4 || 'markdown-header'}">${content}</h4>`;
  });
  markdown = markdown.replace(/^### (.*?)$/gm, (match, content) => {
      const id = generateID(content);
      return `<h3 id="${id}" class="${dynamicClasses.h3 || 'markdown-header'}">${content}</h3>`;
  });
  markdown = markdown.replace(/^## (.*?)$/gm, (match, content) => {
      const id = generateID(content);
      return `<h2 id="${id}" class="${dynamicClasses.h2 || 'markdown-header'}">${content}</h2>`;
  });
  markdown = markdown.replace(/^# (.*?)$/gm, (match, content) => {
      const id = generateID(content);
      return `<h1 id="${id}" class="${dynamicClasses.h1 || 'markdown-header'}">${content}</h1>`;
  });

  // Convert bold (e.g., **bold** or __bold__)
  markdown = markdown.replace(/\*\*(.*?)\*\*/g, `<strong class="${dynamicClasses.bold || 'markdown-bold'}">$1</strong>`);
  markdown = markdown.replace(/__(.*?)__/g, `<strong class="${dynamicClasses.bold || 'markdown-bold'}">$1</strong>`);

  // Convert italic (e.g., *italic* or _italic_)
  markdown = markdown.replace(/\*(.*?)\*/g, `<em class="${dynamicClasses.italic || 'markdown-italic'}">$1</em>`);
  markdown = markdown.replace(/_(.*?)_/g, `<em class="${dynamicClasses.italic || 'markdown-italic'}">$1</em>`);

  // Convert unordered lists (e.g., * item or - item)
  markdown = markdown.replace(/^(\*|\-|\+) (.*?)$/gm, `<ul class="${dynamicClasses.ul || 'markdown-list'}"><li>$2</li></ul>`);

  // Convert ordered lists (e.g., 1. item)
  markdown = markdown.replace(/^(\d+)\. (.*?)$/gm, `<ol class="${dynamicClasses.ol || 'markdown-list'}"><li>$2</li></ol>`);

  // Convert links (e.g., [link text](http://url))
  markdown = markdown.replace(/\[([^\]]+)\]\(([^)]+)\)/g, `<a href="$2" class="${dynamicClasses.link || 'markdown-link'}">$1</a>`);

  // Convert images (e.g., ![alt text](http://url))
  markdown = markdown.replace(/!\[([^\]]+)\]\(([^)]+)\)/g, `<img src="$2" alt="$1" class="${dynamicClasses.img || 'markdown-image'}">`);

  // Convert blockquotes (e.g., > quote)
  markdown = markdown.replace(/^> (.*?)$/gm, `<blockquote class="${dynamicClasses.blockquote || 'markdown-blockquote'}">$1</blockquote>`);

  // Convert horizontal rules (e.g., --- or ***)
  markdown = markdown.replace(/(\*\*\*|\-\-\-|\_\_\_)/g, `<hr class="${dynamicClasses.hr || 'markdown-hr'}">`);

  // Convert line breaks (e.g., two spaces at the end of a line)
  markdown = markdown.replace(/\n\s*\n/g, `<br class="${dynamicClasses.br || 'markdown-br'}">`);

  return `<div class="${dynamicClasses.wrapper || 'markdown-content'}">${markdown}</div>`;
}

function htmlToMarkdown(html, options = {}) {
  // Default options
  const {
      allowedTags = ['b', 'i', 'strong', 'em', 'a', 'img', 'ul', 'ol', 'li', 'p', 'blockquote', 'pre', 'code', 'br', 'hr'],
      dangerousProtocols = /(javascript:|data:|vbscript:|file:)/i,
      escapeSpecialChars = true
  } = options;

  // Helper function to sanitize HTML by removing unsafe attributes
  function sanitizeHTML(html) {
      // Create a temporary div element to parse HTML
      const doc = new DOMParser().parseFromString(html, 'text/html');
      const elements = doc.body.getElementsByTagName('*');

      // Sanitize each element by checking its tag and removing dangerous attributes
      for (let element of elements) {
          // Remove all attributes that start with 'on' (like onclick, onmouseover, etc.)
          for (let attr of element.attributes) {
              if (attr.name.startsWith('on') || attr.name === 'style' || dangerousProtocols.test(attr.value)) {
                  element.removeAttribute(attr.name);
              }
          }

          // Remove any elements that are not in the allowedTags list
          if (!allowedTags.includes(element.tagName.toLowerCase())) {
              element.remove();
          }
      }

      // Return sanitized HTML as a string
      return doc.body.innerHTML;
  }

  // Sanitize the input HTML before converting
  const sanitizedHTML = sanitizeHTML(html);

  // Convert sanitized HTML to Markdown
  let markdown = sanitizedHTML;

  // Convert <b> and <strong> to **bold**
  markdown = markdown.replace(/<(b|strong)>(.*?)<\/\1>/g, '**$2**');

  // Convert <i> and <em> to *italic*
  markdown = markdown.replace(/<(i|em)>(.*?)<\/\1>/g, '*$2*');

  // Convert <a href="..."> to [link](...)
  markdown = markdown.replace(/<a href="([^"]+)">([^<]+)<\/a>/g, '[$2]($1)');

  // Convert <img src="..."> to ![alt](...)
  markdown = markdown.replace(/<img src="([^"]+)" alt="([^"]*)"/g, '![alt]($1)');

  // Convert <ul> and <ol> to lists
  markdown = markdown.replace(/<ul>(.*?)<\/ul>/gs, (match, p1) => {
      return p1.replace(/<li>(.*?)<\/li>/g, '- $1');
  });

  markdown = markdown.replace(/<ol>(.*?)<\/ol>/gs, (match, p1) => {
      return p1.replace(/<li>(.*?)<\/li>/g, (match, p1, index) => `${index + 1}. ${p1}`);
  });

  // Convert <p> to paragraphs
  markdown = markdown.replace(/<p>(.*?)<\/p>/g, '$1\n\n');

  // Convert <blockquote> to blockquotes
  markdown = markdown.replace(/<blockquote>(.*?)<\/blockquote>/gs, '> $1');

  // Convert <pre><code> to code blocks
  markdown = markdown.replace(/<pre><code>(.*?)<\/code><\/pre>/gs, '```\n$1\n```');

  // Convert <br> to line breaks
  markdown = markdown.replace(/<br\s*\/?>/g, '\n');

  // Convert <hr> to horizontal rules
  markdown = markdown.replace(/<hr\s*\/?>/g, '\n---\n');

  // Optionally escape special Markdown characters if required
  if (escapeSpecialChars) {
      markdown = markdown.replace(/([\\`*{}\[\]()#+\-.!_>])/g, '\\$1');
  }

  // Return the final sanitized Markdown content
  return markdown.trim();
}
