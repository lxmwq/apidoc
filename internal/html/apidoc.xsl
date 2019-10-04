<?xml version="1.0" encoding="UTF-8"?>

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:output
    method="html"
    encoding="utf-8"
    indent="yes"
    version="5.0"
    doctype-system="about:legacy-compat" />

<!-- 替换字符串中特定的字符 -->
<xsl:template name="replace">
    <xsl:param name="text" />
    <xsl:param name="old" />
    <xsl:param name="new" />
    <xsl:choose>
        <xsl:when test="contains($text, $old)">
            <xsl:value-of select="substring-before($text, $old)" />
            <xsl:value-of select="$new" />
            <xsl:call-template name="replace">
                <xsl:with-param name="text" select="substring-after($text, $old)" />
                <xsl:with-param name="old" select="$old" />
                <xsl:with-param name="new" select="$new" />
            </xsl:call-template>
        </xsl:when>
        <xsl:otherwise>
            <xsl:value-of select="$text" />
        </xsl:otherwise>
    </xsl:choose>
</xsl:template>

<!-- 根据 method 和 path 生成唯一的 ID -->
<xsl:template name="get-api-id">
    <xsl:param name="method" />
    <xsl:param name="path" />

    <xsl:variable name="p1">
        <xsl:call-template name="replace">
            <xsl:with-param name="text" select="$path" />
            <xsl:with-param name="old" select="'/'" />
            <xsl:with-param name="new" select="'_'" />
        </xsl:call-template>
    </xsl:variable>
    <xsl:variable name="p2">
        <xsl:call-template name="replace">
            <xsl:with-param name="text" select="$p1" />
            <xsl:with-param name="old" select="'{'" />
            <xsl:with-param name="new" select="'-'" />
        </xsl:call-template>
    </xsl:variable>

    <xsl:value-of select="$method" />
    <xsl:call-template name="replace">
        <xsl:with-param name="text" select="$p2" />
        <xsl:with-param name="old" select="'}'" />
        <xsl:with-param name="new" select="'-'" />
    </xsl:call-template>
</xsl:template>

<xsl:template match="/">
    <html>
        <head>
            <title><xsl:value-of select="apidoc/title" /></title>
            <meta charset="UTF-8" />
            <meta name="viewport" content="width=device-width,initial-scale=1,maximum-scale=1"/>
            <link rel="stylesheet" type="text/css" href="./apidoc.css" />
            <link rel="icon" type="image/png" href="./icon.png" />
            <link rel="license">
                <xsl:attribute name="href"><xsl:value-of select="apidoc/license/@url"/></xsl:attribute>
            </link>
            <script src="./apidoc.js"></script>
        </head>
        <body>
            <header>
                <h1>
                    <span class="app-name"><img src="./icon.svg" />apidoc</span>
                </h1>
            </header>

            <aside>
                <h2>Servers</h2>
                <ul class="servers-selector">
                    <xsl:for-each select="apidoc/server">
                    <li>
                        <xsl:attribute name="data-server"><!-- chrome 和 safari 必须要在其它元素之前 -->
                            <xsl:value-of select="@name" />
                        </xsl:attribute>
                        <label><input type="checkbox" /><xsl:value-of select="@name" /></label>
                    </li>
                    </xsl:for-each>
                </ul>

                <h2>Tags</h2>
                <ul class="tags-selector">
                    <xsl:for-each select="apidoc/tag">
                    <li>
                        <xsl:attribute name="data-server">
                            <xsl:value-of select="@name" />
                        </xsl:attribute>
                        <label><input type="checkbox" /><xsl:value-of select="@name" /></label>
                    </li>
                    </xsl:for-each>
                </ul>

                <h2>Methods</h2>
                <ul class="methods-selector">
                    <!-- 浏览器好像都不支持 xpath 2.0，所以无法使用 distinct-values -->
                    <!-- xsl:for-each select="distinct-values(/apidoc/api/@method)" -->
                    <xsl:for-each select="/apidoc/api/@method[not(../preceding-sibling::api/@method = .)]">
                    <li>
                        <xsl:attribute name="data-method">
                            <xsl:value-of select="." />
                        </xsl:attribute>
                        <label><input type="checkbox" /><xsl:value-of select="." /></label>
                    </li>
                    </xsl:for-each>
                </ul>
            </aside>

            <main>
                <div class="content">
                    <xsl:value-of select="/apidoc/content" />
                </div>

                <xsl:for-each select="/apidoc/server">
                <article class="server">
                    <h3>
                        <xsl:value-of select="@url"/>
                        <xsl:value-of select="@name"/>
                    </h3>
                    <div class="summary">
                        <xsl:value-of select="."/>
                    </div>
                </article>
                </xsl:for-each>

                <!-- api -->
                <xsl:for-each select="apidoc/api">
                <article>
                <xsl:attribute name="data-method">
                    <xsl:value-of select="@method" />
                </xsl:attribute>
                <xsl:attribute name="id">
                    <xsl:call-template name="get-api-id">
                        <xsl:with-param name="path" select="path/@path" />
                        <xsl:with-param name="method" select="@method" />
                    </xsl:call-template>
                </xsl:attribute>

                    <h3>
                        <span class="method">
                        <xsl:value-of select="@method" />
                        </span>
                        <xsl:value-of select="path/@path" />
                    </h3>
                    <div class="summary">
                        <xsl:value-of select="@summary" />
                    </div>

                    <div class="body">
                        <div class="request">
                            <h4>请求</h4>
                            <!-- request -->
                        </div>
                        <div class="response">
                            <h4>返回</h4>
                            <!-- response -->
                        </div>
                    </div>
                </article>
                </xsl:for-each>
                <!-- end api -->
            </main>

            <footer>
            <p>文档由 <a href="https://github.com/caixw/apidoc">apidoc</a> 生成！</p>
            </footer>
        </body>
    </html>
</xsl:template>
</xsl:stylesheet>
